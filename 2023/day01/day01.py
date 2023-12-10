from typing import Union


class TrieNode:
    children: dict
    leaf_value: int

    def __init__(self):
        self.children = {}
        self.leaf_value = -1

    def lookup(self, c: str) -> Union['TrieNode', None]:
        return self.children.get(c, None)
    
    def insert(self, word: str, leaf_value: int):
        c, *rest = word
        if c not in self.children:
            self.children[c] = TrieNode()

        next_node = self.children[c]

        if len(rest) == 0:
            next_node.leaf_value = leaf_value
        else:
            next_node.insert(rest, leaf_value)

    def stringify(self) -> str:
        if self.leaf_value != -1:
            return str(self.leaf_value)

        if len(self.children) == 0:
            return ""

        output = ""
        for c, node in self.children.items():
            output += f"{c}: [{node.stringify()}]"

        return output


trie = TrieNode()


def init_number_trie():
    trie.insert("zero", 0)
    trie.insert("one", 1)
    trie.insert("two", 2)
    trie.insert("three", 3)
    trie.insert("four", 4)
    trie.insert("five", 5)
    trie.insert("six", 6)
    trie.insert("seven", 7)
    trie.insert("eight", 8)
    trie.insert("nine", 9)

    print(trie.stringify())


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    init_number_trie()

    first = last = -1
    sum = 0
    curr_nodes: list(TrieNode) = []
    for line in lines:
        for c in line.lower():
            if c.isdigit():
                last = ord(c) - ord('0')
                if first == -1:
                    first = last

            elif c.isalnum():
                next_nodes = []
                curr_nodes.append(trie)
                for node in curr_nodes:
                    next_node = node.lookup(c)
                    if next_node is not None:
                        val = next_node.leaf_value
                        if val == -1:
                            next_nodes.append(next_node)
                        else:
                            last = val
                            if first == -1:
                                first = last

                curr_nodes = next_nodes

        if first < 0:
            continue

        curr_value = 10 * first + last
        print(curr_value)
        sum += curr_value
        first = -1
        curr_nodes = []

    print(sum)


if __name__ == "__main__":
    main()