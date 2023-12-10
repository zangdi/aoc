from collections import defaultdict


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    sum = 0
    for line in lines:
        line = line[5:-1]
        colon_separated = line.split(": ")
        ball_sets = colon_separated[1].split("; ")
        curr_max = defaultdict(int)
        for ball_set in ball_sets:
            balls = ball_set.split(", ")
            for ball in balls:
                amount_str, colour = ball.split(" ")
                amount = int(amount_str)
                if amount > curr_max[colour]:
                    curr_max[colour] = amount

        power = curr_max["red"] * curr_max["green"] * curr_max["blue"]
        print(f"curr_max: {curr_max}, power: {power}")
        sum += power

    print(sum)


if __name__ == "__main__":
    main()