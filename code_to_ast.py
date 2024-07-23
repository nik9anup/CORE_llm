import json

def node_to_json(node, source_code):
    start_byte = node.start_byte
    end_byte = node.end_byte
    children = [node_to_json(child, source_code) for child in node.children] if node.children else []
    json_node = {
        "type": node.type,
        #"line_number": node.start_point.row + 1,
        #"start_byte": start_byte,
        #"end_byte": end_byte,
        "text": source_code[start_byte:end_byte],
        "children": children
    }
    return json_node

def get_ast(go_code_in, parser):
    go_code_bytes = go_code_in.encode('utf-8')
    tree = parser.parse(go_code_bytes)
    root_node = tree.root_node
    root_json = node_to_json(root_node, go_code_in)
    return root_json
    """
    filename = "ast.json"
    with open(filename, "w") as file:
        json.dump(root_json, file, indent=2)
    print(f"AST JSON saved to {filename}")
    """