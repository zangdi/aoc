from collections import defaultdict


def check_and_get_gears(lines: list, x: int, y: int) -> any:
    digits = []
    for i in range(-1, 2):
        prevNum = False
        for j in range(-1, 2):
            if lines[x + i][y + j].isdigit():
                if not prevNum:
                    digits.append((x + i, y + j))

                prevNum = True
            else:
                prevNum = False

    if len(digits) != 2:
        return False, ((0, 0), (0, 0))
    return True, (digits[0], digits[1])


def get_number(lines: list, gear: tuple) -> int:
    idx = gear[1]
    while idx >= 1 and lines[gear[0]][idx - 1].isdigit():
        idx -= 1

    number = 0
    while idx < 140 and lines[gear[0]][idx].isdigit():
        digit = ord(lines[gear[0]][idx]) - ord("0")
        number = number * 10 + digit
        idx += 1

    print(number)
    return number


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    sum = 0
    for x, line in enumerate(lines):
        for y, ch in enumerate(line):
            if ch != "*":
                continue

            is_gear, gears = check_and_get_gears(lines, x, y)
            if not is_gear:
                continue

            print(f"got gear at {gears}")
            ratio = get_number(lines, gears[0]) * get_number(lines, gears[1])
            sum += ratio

    print(sum)


if __name__ == "__main__":
    main()