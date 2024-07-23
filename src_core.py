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

# Set environment variables
os.environ["ANTHROPIC_API_KEY"] = os.getenv("ANTHROPIC_API_KEY")
os.environ["GOOGLE_API_KEY"] = os.getenv("GOOGLE_API_KEY")

# Get GitHub token and PR number
github_token = os.getenv("GITHUB_TOKEN")
pr_number = int(os.getenv("PR_NUMBER"))

# Define GitHub client
g = Github(github_token)
repo = g.get_repo("owner/repo")  # replace with your repo

# Get the PR data
pr = repo.get_pull(pr_number)

# Get the files changed in the PR
files = pr.get_files()
go_code_in = ""
for file in files:
    if file.filename.endswith(".go"):
        go_code_in += file.patch

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
You are a code refactoring assistant specialized in Golang. Your task is to review the provided code and suggest using appropriate functions from the language's standard library or popular packages to make the code more concise and efficient, while maintaining its functionality.

Do not include any reasoning, comments or text in the output except the code.
Do not include a main function.
Kindly use package main BEFORE the imported packages.

Here is an example for your reference in which the manual parsing and validation of JSON is made concise using json.Unmarshal() function from encoding/json library while maintaining functionality and semantics:

Example Input:
package main

import (
    "fmt"
    "errors"
    "strings"
)

func parseAndValidateJSON(jsonStr string) (map[string]string, error) {{
    jsonStr = removeWhitespace(jsonStr)

    if !startsWith(jsonStr, "{{") || !endsWith(jsonStr, "}}") {{
        return nil, errors.New("invalid JSON format")
    }}

    jsonStr = removeBraces(jsonStr)

    pairs := splitIntoPairs(jsonStr)
    result := make(map[string]string)

    for _, pair := range pairs {{
        kv := splitKeyValue(pair)
        if len(kv) != 2 {{
            return nil, errors.New("invalid key-value pair")
        }}

        key := removeQuotes(kv[0])
        value := removeQuotes(kv[1])

        if key == "" {{
            return nil, errors.New("invalid key")
        }}

        result[key] = value
    }}

    return result, nil
}}

// Helper functions (simulated std library functions)
func removeWhitespace(str string) string {{
    return strings.ReplaceAll(str, " ", "")
}}

func startsWith(str, prefix string) bool {{
    return strings.HasPrefix(str, prefix)
}}

func endsWith(str, suffix string) bool {{
    return strings.HasSuffix(str, suffix)
}}

func removeBraces(str string) string {{
    return str[1 : len(str)-1]
}}

func splitIntoPairs(str string) []string {{
    return strings.Split(str, ",")
}}

func splitKeyValue(str string) []string {{
    return strings.Split(str, ":")
}}

func removeQuotes(str string) string {{
    return strings.Trim(str, "\"")
}}

Example Output:
package main

import (
    "fmt"
    "encoding/json"
)

func parseAndValidateJSON(jsonStr string) (map[string]string, error) {{
    var data map[string]string
    if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {{
        return nil, errors.New("invalid JSON format")
    }}
    return data, nil
}}

input code snippet: 
{query_str}
"""

text_qa_template = PromptTemplate(text_qa_template_str)

# Extract relevant Go code parts
query_str = ""
go_codes_out = extract_relevant_part.get_relevant_part(go_code_in, parser)
for go_code_out in go_codes_out:
    query_str += go_code_out + "\n\n"

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

# Query the engine with the extracted code
response = query_engine.query(query_str)
print(response.response)
