import itertools
import math
import re

from collections import defaultdict


class graphNode:
    name: str
    left: 'graphNode'
    right: 'graphNode'

    def __init__(self, name):
        self.name = name
        self.left = None
        self.right = None


def getOrCreate(name: str, node_map: dict):
    if name not in node_map:
        node_map[name] = graphNode(name)

    return node_map[name]


class nodeIterator:
    loopStart: int
    lengths: int
    idx: int
    cycleLength: int

    def __init__(self, loopStart, lengths, cycleLength) -> None:
        self.loopStart = loopStart
        self.lengths = lengths
        self.idx = 0
        self.cycleLength = cycleLength

    def getValue(self) -> int:
        return self.lengths[self.idx]
    
    def increment(self):
        prevIdx = self.idx
        self.idx = (self.idx + 1) % len(self.lengths)
        self.lengths[prevIdx] += self.cycleLength

    def __repr__(self) -> str:
        return "{" + f"loopstart: {self.loopStart}, lengths: {self.lengths}, index: {self.idx}, cyclelength: {self.cycleLength}" + "}"


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    instructions = lines[0].strip()
    node_map = defaultdict(lambda: None)
    for line in lines[2:]:
        node_name, left_name, right_name = re.findall(r"\w+", line)
        node = getOrCreate(node_name, node_map)
        left = getOrCreate(left_name, node_map)
        right = getOrCreate(right_name, node_map)

        node.left = left
        node.right = right

    node_to_lengths = {}
    just_lengths = []
    for node in node_map.values():
        if node.name[2] != "A":
            continue

        lengths = []
        visited = set()
        curr = node
        steps = 0
        print(f"for {curr.name}")
        cycleLength = 0
        while curr is not None:
            idx = steps % len(instructions)
            visit = (curr.name, idx)
            if visit in visited:
                cycleLength = steps - idx
                print(f"collision at {visit} after {steps}")
                break

            visited.add(visit)
            steps += 1

            if instructions[idx] == "L":
                curr = curr.left
            else:
                curr = curr.right

            if curr.name[2] == "Z":
                print(curr.name, steps)
                lengths.append(steps)
                just_lengths.append(steps)

        print(sorted(visited))
        node_to_lengths[node.name] = nodeIterator(idx, lengths, cycleLength)

    print(just_lengths)
    print(math.lcm(*just_lengths))

    # print(node_to_lengths)
    # curr_max = 0
    # while any([it.getValue() != curr_max for it in node_to_lengths.values()]) and curr_max < 17163415311930:
    #     for name, it in node_to_lengths.items():
    #         value = it.getValue()
    #         while value < curr_max:
    #             it.increment()
    #             value = it.getValue()
    #             # print(f"{name}: value = {value}")

    #         curr_max = value
    #         # print(f"{name} -> currmax: {curr_max}")

    # print(curr_max)



    # steps = 0
    # currs = [node for node in node_map.values() if node.name[2] == "A"]
    # while any([True for node in currs if node.name[2] != "Z"]):
    #     idx = steps % len(instructions)
    #     steps += 1

    #     if instructions[idx] == "L":
    #         currs = [curr.left for curr in currs]
    #     else:
    #         currs = [curr.right for curr in currs]

    # print(steps)


if __name__ == "__main__":
    main()