import sys

from dataclasses import dataclass
from enum import IntEnum

VERTICAL_SNEAK = ["JL", "7F", "||", "J|", "|L", "7|", "|F"]
HORIZONTAL_SNEAK = ["J7", "LF", "--", "J-", "-7", "L-", "-F"]

VERTICAL_STOPPER = ["L"]


@dataclass(frozen=True)
class Coordinates:
    x: int
    y: int


class Direction(IntEnum):
    EAST = 2
    NORTH = 1
    NONE = 0
    SOUTH = -1
    WEST = -2

    def opposite(self) -> 'Direction':
        return Direction(-self)
    
    def outDirection(self) -> str:
        if self == Direction.NORTH:
            return "SJL|"
        elif self == Direction.SOUTH:
            return "S7F|"
        elif self == Direction.EAST:
            return "SJ7-"
        elif self == Direction.WEST:
            return "SLF-"
        
    def inDirection(self) -> str:
        return self.opposite().outDirection()
    
    def getCoordinates(self, old: Coordinates) -> Coordinates:
        if self == Direction.NORTH:
            return Coordinates(old.x - 1, old.y)
        elif self == Direction.SOUTH:
            return Coordinates(old.x + 1, old.y)
        elif self == Direction.EAST:
            return Coordinates(old.x, old.y - 1)
        elif self == Direction.WEST:
            return Coordinates(old.x, old.y + 1)


def dfs(graph: list, newGraph: list, pos: Coordinates, visited: set, d: Direction) -> tuple:
    if pos.x < 0 or pos.x >= len(graph) or pos.y < 0 or pos.y >= len(graph[0]):
        return False

    if d != Direction.NONE and graph[pos.x][pos.y] not in d.inDirection():
        return False
    
    if pos in visited:
        if graph[pos.x][pos.y] == "S":
            newGraph[(pos.x * 3) + 1][(pos.y * 3) + 1] = 1
            if d == Direction.NORTH:
                newGraph[(pos.x * 3) + 2][(pos.x * 3) + 1] = 1
            elif d == Direction.SOUTH:
                newGraph[(pos.x * 3)][(pos.x * 3) + 1] = 1
            elif d == Direction.EAST:
                newGraph[(pos.x * 3) + 1][(pos.x * 3) + 2] = 1
            elif d == Direction.WEST:
                newGraph[(pos.x * 3) + 1][(pos.x * 3)] = 1
            return True

        return False
    
    visited.add(pos)
    found = False
    for direction in [Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST]:
        if graph[pos.x][pos.y] in direction.outDirection() and d != direction.opposite():
            found = dfs(graph, newGraph, direction.getCoordinates(pos), visited, direction)
            if found:
                newGraph[pos.x][pos.y] = 1
                return True


def fill(graph: list, pos: Coordinates):
    if pos.x < 0 or pos.x >= len(graph) or pos.y < 0 or pos.y >= len(graph[0]) or graph[pos.x][pos.y] != 0:
        return
    
    graph[pos.x][pos.y] = 2
    fill(graph, Coordinates(pos.x - 1, pos.y))
    fill(graph, Coordinates(pos.x + 1, pos.y))
    fill(graph, Coordinates(pos.x, pos.y - 1))
    fill(graph, Coordinates(pos.x, pos.y + 1))


def isWet(graph: list, x: int, y: int):
    for i in range(2):
        for j in range(2):
            if graph[3 * x + i][3 * y + j] != 0:
                return True

    return False


def compact(graph: list) -> list:
    newGraph = []
    for i in range(len(graph) // 3):
        row = []
        for j in range(len(graph[i]) // 3):
            if isWet(graph, i, j):
                row.append(0)
            else:
                row.append(1)

        newGraph.append(row)

    return newGraph


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    start = Coordinates(0, 0)
    newGraph = []
    for x, line in enumerate(lines):
        newGraph.append([0 for _ in range(len(line) * 3)])
        newGraph.append([0 for _ in range(len(line) * 3)])
        newGraph.append([0 for _ in range(len(line) * 3)])

        y = line.find("S")
        if y != -1:
            start = Coordinates(x, y)
    

    found = dfs(lines, newGraph, start, set(), Direction.NONE)
    fill(newGraph, Coordinates(0, 0))

    print("\n".join(["".join([str(x) for x in line]) for line in newGraph]))
    graph2 = compact(newGraph)

    print("\n".join(["".join([str(x) for x in line]) for line in graph2]))

    print(sum([num for line in graph2 for num in line]))


if __name__ == "__main__":
    sys.setrecursionlimit(800000)
    main()