import re

def allSame(arr: list) -> bool:
    if len(arr) == 0:
        return True
    
    return all([num == arr[0] for num in arr[1:]])


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    sum = 0
    for line in lines:
        nums = [int(x) for x in re.findall(r"-?\d+", line)]

        diffs = nums
        firsts = [nums[0]]
        while not allSame(diffs):
            prev = diffs
            diffs = []
            for i in range(len(prev) - 1):
                diffs.append(prev[i + 1] - prev[i])

            print(diffs)
            firsts.append(diffs[0])

        curr = 0
        for num in reversed(firsts):
            curr = num - curr

        print(curr)
        sum += curr

    print(sum)


if __name__ == "__main__":
    main()