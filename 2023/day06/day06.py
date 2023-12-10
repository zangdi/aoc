import re
import sys


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    counts = [1 for _ in lines]
    r = re.compile(r"\d+")
    for i, line in enumerate(lines):
        line = line[8:]
        winnings, card = line.split("|")

        winningSet = set(re.findall(r, winnings))
        matches = len([1 for x in re.findall(r, card) if x in winningSet])

        for j in range(i + 1, i + matches + 1):
            counts[j] += counts[i]

    print(sum(counts))


if __name__ == "__main__":
    main()