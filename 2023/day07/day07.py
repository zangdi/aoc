from collections import defaultdict


def str_to_val(s: str) -> int:
    if s == "A":
        return 13
    elif s == "K":
        return 12
    elif s == "Q":
        return 11
    elif s == "T":
        return 10
    elif s == "J":
        return 1
    else:
        return int(s)


class Wager:
    card: str
    value: int
    strength: int

    def __init__(self, card, value, strength) -> None:
        self.card = card
        self.value = value
        self.strength = strength

    def __lt__(self, other) -> bool:
        if self.strength != other.strength:
            return self.strength < other.strength
        
        for i in range(5):
            if self.card[i] != other.card[i]:
                return str_to_val(self.card[i]) < str_to_val(other.card[i])
            
        return True
    
    def __repr__(self) -> str:
        return "{" + f"card: {self.card}, value: {self.value}, strength: {self.strength}" + "}"


def main():
    with open("input.txt") as f:
        lines = f.readlines()

    wagers = []
    for line in lines:
        card, bet = line.split(" ")
        value = int(bet)

        counts = defaultdict(int)
        for c in card:
            counts[c] += 1

        jokers = counts.pop("J") if "J" in counts else 0
        if jokers == 5:
            wagers.append(Wager(card, value, 6))
            continue

        sets = defaultdict(int)
        for v in counts.values():
            sets[v] += 1

        print(sets)
        most = sorted(sets, reverse=True)[0]
        sets[most] -= 1
        sets[most + jokers] += 1

        strength = 0
        if sets[5] == 1:
            strength = 6
        elif sets[4] == 1:
            strength = 5
        elif sets[3] == 1 and sets[2] == 1:
            strength = 4
        elif sets[3] == 1:
            strength = 3
        elif sets[2] == 2:
            strength = 2
        elif sets[2] == 1:
            strength = 1

        wagers.append(Wager(card, value, strength))

    wagers.sort()

    print(wagers)

    total = 0
    for i, wager in enumerate(wagers, start = 1):
        total += i * wager.value

    print(total)


if __name__ == "__main__":
    main()