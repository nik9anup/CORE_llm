import os
import chromadb
from github import Github
from llama_index.embeddings.gemini import GeminiEmbedding
from llama_index.core import Settings, VectorStoreIndex, SimpleDirectoryReader, StorageContext
from llama_index.llms.gemini import Gemini
from llama_index.llms.anthropic import Anthropic
from llama_index.vector_stores.chroma import ChromaVectorStore
from llama_index.core.retrievers import VectorIndexRetriever
from llama_index.core.node_parser import CodeSplitter
import tree_sitter_go as tsgo
from tree_sitter import Language, Parser
from llama_index.core import PromptTemplate
import extract_relevant_part
import pysqlite3
import sys
sys.modules["sqlite3"] = sys.modules.pop("pysqlite3")

# Set environment variables
os.environ["ANTHROPIC_API_KEY"] = os.getenv("ANTHROPIC_API_KEY")
os.environ["GOOGLE_API_KEY"] = os.getenv("GOOGLE_API_KEY")

# Get GitHub token and PR number
github_token = os.getenv("GITHUB_TOKEN")
pr_number = int(os.getenv("PR_NUMBER"))
#repo_owner = os.getenv("REPO_OWNER")
repo_name = os.getenv("REPO_NAME")

# Define GitHub client
g = Github(github_token)

repo = g.get_repo(repo_name)

# Get the PR data
pr = repo.get_pull(pr_number)

# Get the files changed in the PR
go_code_in= " "

commits = pr.get_commits()

for commit in commits:
    files = commit.files
    for file in files:
      filename = file.filename
      if filename.endswith('.go'):
         contents = repo.get_contents(filename, ref=commit.sha).decoded_content
         go_code_in += contents

print(go_code_in)

# Initialize ChromaDB client and collection
chroma_client = chromadb.PersistentClient(path="./main")
chroma_collection = chroma_client.get_or_create_collection("astchroma")
documents = SimpleDirectoryReader('data').load_data()

# Define the Go language and parser
GO_LANGUAGE = Language(tsgo.language())
parser = Parser()
parser.set_language(GO_LANGUAGE)

# Initialize CodeSplitter
splitter = CodeSplitter.from_defaults(language='go', parser=parser)

# Define the prompt template
text_qa_template_str = """
You are a code refactoring assistant specialised in golang. Your task is to review the provided code and suggest using appropriate functions from the language's standard library or popular packages to make the code more concise and efficient, while maintaining its functionality.\n
Do not include any reasoning, comments or text in the output except the code.\n
Do not add any other function that isn't there in the input. \n
Do not remove any lines of code from user input unless replacing it with a library function. \n
Kindly use package main BEFORE the imported packages. \n


input code snippet: \n
{query_str}



"""

text_qa_template = PromptTemplate(text_qa_template_str)

# Extract relevant Go code parts
query_str = ""
go_codes_out = extract_relevant_part.get_relevant_part(go_code_in, parser)

# Initialize tokenizer and models
tokenizer = Anthropic().tokenizer
llm = Anthropic(model="claude-3-haiku-20240307")
embed_model = GeminiEmbedding(model_name="models/text-embedding-004")

# Initialize vector store and index
vector_store = ChromaVectorStore(chroma_collection=chroma_collection)
storage_context = StorageContext.from_defaults(vector_store=vector_store)
index = VectorStoreIndex.from_vector_store(vector_store, embed_model=embed_model, transformations=[splitter])

# Initialize query engine
query_engine = index.as_query_engine(text_qa_template=text_qa_template, llm=llm)

result = []
for go_code_out in go_codes_out:
    query_str = "".join(go_code_out)
    response=query_engine.query(query_str)
    result.append(response.response)

text_qa_template_str_1 = """Make one meaningful and concise go language snippet by combining all the code snippets you have received in the query. \n\n
Do not include any reasoning, comments or text in the output except the code.\n  
Do not remove any lines of code. \n

input code list: \n
{query_str}
"""

text_qa_template_1 = PromptTemplate(text_qa_template_str_1)

query_engine_1 = index.as_query_engine(text_qa_template=text_qa_template_1,llm=llm)

query_str_1 = str(result)
response_1 = query_engine_1.query(query_str_1)
print(response_1.response)
