import pdb
citations = [4, 3, 2, 4, 5]
N = 5
cases  = sorted(list(set(citations)), reverse=True)

print(cases)

for case in cases:
    bla = sum([1 for x in citations if x >= case])
    if bla >= case:
        print(case)
        break