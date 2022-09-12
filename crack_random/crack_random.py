# // a>>b = a/(2^b)
# // a<<b = a*(2^b)

def nxt(seed):
    seed2 = (seed * 25214903917 + 11) % (2**48)
    return seed2 >> 16


n1 = int(input("n1: "))
n2 = int(input("n2: "))

for i in range(65536):
    seed = (n1 * 65536) + i
    if nxt(seed) == n2:
        print(">>> Found seed", seed)
        break

for i in range(5):
    x = seed >> 16
    if x >= 2**31:
        x -= 2**32
    print(x)
    seed = (seed * 25214903917 + 11) % (2**48)