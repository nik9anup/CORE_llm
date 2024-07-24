"""need to remove lines with ### for approach 1"""
import random

"""
def get_line_node(node, line_no):
    if node.type != "source_file" and (node.start_point.row + 1) == line_no :
        return node
    elif node.children:
        for child_node in node.children:
            ic_node = get_line_node(child_node, line_no)
            if ic_node:
                return ic_node
    else:
        return None

def get_initial_context(root_node, comment_on_line_no, go_code):
    initial_context_node = get_line_node(root_node, comment_on_line_no)
    if initial_context_node is not None:
        while initial_context_node.parent.type != "source_file":
            initial_context_node = initial_context_node.parent
        return go_code[initial_context_node.start_byte : initial_context_node.end_byte]
    else:
        return ""
"""

def search_children(node, identifier, source_code, function, func_nodes):
    if node.children:
        for child_node in node.children:
            found_context = function(child_node, identifier, source_code, func_nodes)
            if found_context[0]:
                return found_context
    #root_node = get_root_node(source_code, parser)
    return "", source_code, func_nodes

def find_child_node(node, node_type):
    for child_node in node.children:
        if child_node.type == node_type:
            return child_node
    return None

def check_type(node, identifier, source_code, func_nodes):
    if node.type == "type_declaration":
        child_node = find_child_node(node, "type_spec")
        if child_node:
            grand_child_node = find_child_node(child_node, "type_identifier")
            if grand_child_node and identifier == source_code[grand_child_node.start_byte : grand_child_node.end_byte]:
                found_context = source_code[node.start_byte : node.end_byte]
                func_nodes = remove_func_node(func_nodes, node)
                #source_code = source_code[:node.start_byte] + source_code[node.end_byte:] ###
                #root_node = get_root_node(source_code, parser)
                return found_context, source_code, func_nodes
    return search_children(node, identifier, source_code, check_type, func_nodes)


def check_type_alias(node, identifier, source_code, func_nodes):
    if node.type == "type_declaration":
        child_node = find_child_node(node, "type_alias")
        if child_node:
            grand_child_node = find_child_node(child_node, "type_identifier")
            if grand_child_node and identifier == source_code[grand_child_node.start_byte : grand_child_node.end_byte]:
                found_context = source_code[node.start_byte : node.end_byte]
                func_nodes = remove_func_node(func_nodes, node)
                #source_code = source_code[:node.start_byte] + source_code[node.end_byte:] ###
                #root_node = get_root_node(source_code, parser)
                return found_context, source_code, func_nodes
    return search_children(node, identifier, source_code, check_type_alias, func_nodes)


def check_const(node, identifier, source_code, func_nodes):
    if node.type == "const_declaration":
        child_node = find_child_node(node, "const_spec")
        if child_node:
            grand_child_node = find_child_node(child_node, "type_identifier")
            if grand_child_node and identifier == source_code[grand_child_node.start_byte : grand_child_node.end_byte]:
                found_context = source_code[node.start_byte : node.end_byte]
                func_nodes = remove_func_node(func_nodes, node)
                #source_code = source_code[:node.start_byte] + source_code[node.end_byte:] ###
                #root_node = get_root_node(source_code, parser)
                return found_context, source_code, func_nodes
    return search_children(node, identifier, source_code, check_const, func_nodes)


def check_function(node, identifier, source_code, func_nodes):
    if node.type == "function_declaration":
        child_node = find_child_node(node, "identifier")
        if child_node and identifier == source_code[child_node.start_byte : child_node.end_byte]:
            found_context = source_code[node.start_byte : node.end_byte]
            func_nodes = remove_func_node(func_nodes, node)
            #source_code = source_code[:node.start_byte] + source_code[node.end_byte:] ###
            #root_node = get_root_node(source_code, parser)
            return found_context, source_code, func_nodes
    return search_children(node, identifier, source_code, check_function, func_nodes)

def check_method(node, identifier, source_code, func_nodes):
    if node.type == "method_declaration":
        child_node = find_child_node(node, "field_identifier")
        if child_node and identifier == source_code[child_node.start_byte : child_node.end_byte]:
            found_context = source_code[node.start_byte : node.end_byte]
            func_nodes = remove_func_node(func_nodes, node)
            #source_code = source_code[:node.start_byte] + source_code[node.end_byte:] ###
            #root_node = get_root_node(source_code, parser)
            return found_context, source_code, func_nodes
    return search_children(node, identifier, source_code, check_method, func_nodes)


def check_uniqueness(found_context, context):
    if found_context not in context:
        return '\n' + found_context
    return ""


def get_relevant_context(context, root_node, node, source_code, func_nodes):
    node_type = node.type
    ctx = "".join(context)
    if node_type == "identifier":
        found_context, source_code, func_nodes = check_function(root_node, ctx[node.start_byte : node.end_byte], source_code, func_nodes)
        unique_context = check_uniqueness(found_context, ctx)
        if unique_context:
            context.append(unique_context)
    elif node_type == "field_identifier":
        found_context, source_code, func_nodes = check_method(root_node, ctx[node.start_byte : node.end_byte], source_code, func_nodes)
        unique_context = check_uniqueness(found_context, ctx)
        if unique_context:
            context.append(unique_context)
    elif node_type == "type_identifier":
        found_context, source_code, func_nodes = check_type(root_node, ctx[node.start_byte : node.end_byte], source_code, func_nodes)
        unique_context = check_uniqueness(found_context, ctx)
        if unique_context:
            context.append(unique_context)
        found_context, source_code, func_nodes = check_type_alias(root_node, ctx[node.start_byte : node.end_byte], source_code, func_nodes)
        unique_context = check_uniqueness(found_context, ctx)
        if unique_context:
            context.append(unique_context)
        found_context, source_code, func_nodes = check_const(root_node, ctx[node.start_byte : node.end_byte], source_code, func_nodes)
        unique_context = check_uniqueness(found_context, ctx)
        if unique_context:
            context.append(unique_context)
    #elif node_type == "import_declaration" or node_type == "package_clause":
        #context = context[node.start_byte : node.end_byte]
    elif node.children:
        for child_node in node.children:
            context, source_code, func_nodes = get_relevant_context(context, root_node, child_node, source_code, func_nodes)
    return context, source_code, func_nodes


def append_relevant_context(context, root_node, source_code, parser, func_nodes):
    old_context = []
    new_context = context
    while new_context != old_context:
        rel_node = get_root_node(new_context, parser)
        old_context = new_context
        new_context, source_code, func_nodes = get_relevant_context(new_context, root_node, rel_node, source_code, func_nodes)
    return new_context, source_code, func_nodes

def append_package_import(context, root_node, source_code):
    package_node = find_child_node(root_node, "package_clause")
    import_node = find_child_node(root_node, "import_declaration")
    if import_node:
        context.insert(0, source_code[import_node.start_byte : import_node.end_byte] + '\n')
    if package_node:
        context.insert(0, source_code[package_node.start_byte : package_node.end_byte] + '\n')
    return context

def get_func_nodes(root_node):
    func_nodes = []
    for node in root_node.children:
        if node.type == "function_declaration":
            func_nodes.append(node)
        elif node.type == "method_declaration":
            func_nodes.append(node)
    return func_nodes

def get_root_node(go_code_list, parser):
    go_code = "".join(go_code_list)
    go_code_bytes = go_code.encode('utf-8')
    tree = parser.parse(go_code_bytes)
    root_node = tree.root_node
    return root_node

def remove_func_node(func_nodes, node):
    new_func_nodes = []
    for func_node in func_nodes:
        if func_node.start_byte != node.start_byte:
            new_func_nodes.append(func_node)
    return new_func_nodes

def check_main_func(node, go_code_in):
    if node.type == "function_declaration":
        child_node = find_child_node(node, "identifier")
        if child_node and "main" == go_code_in[child_node.start_byte : child_node.end_byte]:
            return True
    return False



def get_relevant_part(go_code_in, parser):
    root_node = get_root_node(go_code_in, parser)
    func_nodes = get_func_nodes(root_node)
    contexts = []
    final_contexts = []
    while func_nodes:
        #random.shuffle(func_nodes)
        node = func_nodes[0]
        context = [go_code_in[node.start_byte : node.end_byte]]
        func_nodes = remove_func_node(func_nodes, node)
        if check_main_func(node, go_code_in):
            contexts.append(context)
            continue
        #go_code_in = go_code_in[:node.start_byte] + go_code_in[node.end_byte:] ###
        #root_node = get_root_node(go_code_in, parser)
        '''context += get_initial_context(root_node, node, go_code_in)'''
        context, go_code_in, func_nodes = append_relevant_context(context, root_node, go_code_in, parser, func_nodes)
        #func_nodes = get_func_nodes(root_node)
        contexts.append(context)
    for k in range(len(contexts)):
        for l in range(len(contexts[k])):
            contexts[k][l] = contexts[k][l].strip()
    for i in range(len(contexts)):
        for j in range(len(contexts)):
            if i != j and (all(x in contexts[j] for x in contexts[i])):
                break
        else:
            final_contexts.append(contexts[i])
    final_contexts.sort(key = len)
    for index in range(len(final_contexts)):
        final_contexts[index] = append_package_import(final_contexts[index], root_node, go_code_in)
    return final_contexts

"""
import tree_sitter_go as tsgo
from tree_sitter import Language, Parser
# Define the Go language and parser
GO_LANGUAGE = Language(tsgo.language())
parser = Parser(GO_LANGUAGE)
parser.language = GO_LANGUAGE 

with open ("go_code_in.go", "r") as f_in:
    go_code_in = f_in.read()

go_codes_out = get_relevant_part(go_code_in, parser)

for go_code_out in go_codes_out:
    print("\n\n\n\n")
    print(go_code_out)
    '''for i in go_code_out:
        print(i)'''
"""