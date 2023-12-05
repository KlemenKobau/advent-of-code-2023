from typing import List, Tuple


def get_range(ran: str):
    dest_range, source_ran, ran_len = tuple(map(lambda x: int(x), ran.split(" ")))
    delta = dest_range - source_ran

    return (source_ran, source_ran + ran_len), delta     


def get_seed_groups(seed_ranges: List[int]) -> List[Tuple[int, int]]:
    pairs = zip(seed_ranges[::2], seed_ranges[1::2])
    return list(map(lambda x: (x[0], x[0] + x[1]),pairs))


with open("day-5/in.txt") as f:
    content = f.read()

sections = content.split("\n\n")
seed_ranges = list(map(lambda x: int(x), sections[0].split(": ")[1].split(" ")))

seed_groups = get_seed_groups(seed_ranges)

curr_groups = seed_groups

for sec in sections[1:]:
    rows = sec.split("\n")

    new_groups = []

    for curr_start, curr_end in curr_groups:

        handled = False
        for ran in rows[1:]:
            (source_ran, end_ran), delta = get_range(ran)

            if source_ran <= curr_start < end_ran and source_ran < curr_end <= end_ran:
                # all within
                # print(f"({curr_start},{curr_end}) within ({source_ran}, {end_ran})")

                n = (curr_start + delta, curr_end + delta)
                # print(f"({curr_start},{curr_end}) -> {n}")

                new_groups.append(n)
                handled = True
            elif curr_start <= source_ran < curr_end:
                # at the start
                # print(f"({curr_start},{curr_end}) at the start ({source_ran}, {end_ran})")
                n = (source_ran + delta, curr_end + delta)
                # print(f"({source_ran},{curr_end}) -> {n}")

                new_groups.append(n)

                if curr_start != source_ran:
                    curr_groups.append((curr_start, source_ran))
                handled = True

            elif curr_start < end_ran <= curr_end:
                # at the end
                # print(f"({curr_start},{curr_end}) at the end ({source_ran}, {end_ran})")
                n = (curr_start + delta, end_ran + delta)
                new_groups.append(n)

                # print(f"({curr_start},{end_ran}) -> {n}")
                if end_ran != curr_end:
                    curr_groups.append((end_ran, curr_end))
                handled = True
        if not handled:
            new_groups.append((curr_start, curr_end))

    curr_groups = new_groups

# 14149303
print(curr_groups)
print(min(curr_groups, key=lambda x: x[0])[0])
