from llama_index.embeddings.gemini import GeminiEmbedding
from llama_index.core import Settings
from llama_index.core import VectorStoreIndex, SimpleDirectoryReader
from llama_index.llms.gemini import Gemini
from llama_index.llms.anthropic import Anthropic
from llama_index.core import StorageContext
from llama_index.vector_stores.chroma import ChromaVectorStore
from llama_index.core.retrievers import VectorIndexRetriever
from llama_index.core import VectorStoreIndex, SimpleDirectoryReader
from llama_index.core.node_parser import CodeSplitter
from llama_index.core import PromptTemplate
import os
import chromadb
import tree_sitter_go as tsgo
from tree_sitter import Language, Parser
import extract_relevant_part
import code_to_ast

os.environ["ANTHROPIC_API_KEY"] = ""
os.environ["GOOGLE_API_KEY"] = ""

chroma_client = chromadb.PersistentClient(path="./main")
chroma_collection = chroma_client.get_or_create_collection("astchroma")
documents = SimpleDirectoryReader('data').load_data() 

GO_LANGUAGE = Language(tsgo.language())
parser = Parser(GO_LANGUAGE)
parser.language = GO_LANGUAGE 

splitter = CodeSplitter.from_defaults(language='go',parser=parser)

text_qa_template_str = ( 
       """
You are a code refactoring assistant specialised in golang. Your task is to review the provided code and suggest using appropriate functions from the language's standard library or popular packages to make the code more concise and efficient, while maintaining its functionality.\n
Do not include any reasoning, comments or text in the output except the code.\n
Do not add any other function that isn't there in the input. \n
Do not remove any lines of code from user input unless replacing it with a library function. \n
Kindly use package BEFORE the import. \n


input code snippet: \n
{query_str}

"""
     )
text_qa_template = PromptTemplate(text_qa_template_str)

with open ("go_code_in.go", "r") as f_in:
    go_code_in = f_in.read()

go_codes_out = extract_relevant_part.get_relevant_part(go_code_in, parser)
print("Extracted code snippets: \n\n")
for go_code_out in go_codes_out:
    print(go_code_out)
    print("\n\n")

tokenizer = Anthropic().tokenizer

llm = Anthropic(model="claude-3-haiku-20240307")
embed_model = GeminiEmbedding(model_name="models/text-embedding-004")

vector_store = ChromaVectorStore(chroma_collection=chroma_collection)
storage_context = StorageContext.from_defaults(vector_store=vector_store)
index = VectorStoreIndex.from_vector_store(
    vector_store,
    embed_model=embed_model,
    transformations=[splitter]
)

query_engine = index.as_query_engine(text_qa_template=text_qa_template,llm=llm)

result = []
print("Querying the code snippets: \n\n")
for go_code_out in go_codes_out:
    query_str = "".join(go_code_out)
    response=query_engine.query(query_str)
    print(response.response)
    result.append(response.response)
    print("\n\n\n")

text_qa_template_str_1 = """Make one meaningful and concise go language snippet by combining all the code snippets you have received in the query. \n\n
Do not include any reasoning, comments or text in the output except the code.\n  
Do not remove any lines of code. \n
Instead of creating a separate function to demonstrate changes, directly modify the caller function with the replacement code and remove the callee function.\n

input code list: \n
{query_str}
"""

text_qa_template_1 = PromptTemplate(text_qa_template_str_1)

query_engine_1 = index.as_query_engine(text_qa_template=text_qa_template_1,llm=llm)

query_str_1 = str(result)
response = query_engine_1.query(query_str_1)
print("Combined code snippet: \n\n")
print(response.response)

with open("go_code_out.go", "w") as f_out:
    f_out.write(response.response)
