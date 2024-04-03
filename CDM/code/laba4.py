from icecream import ic
ic.configureOutput(prefix='')

import sympy as sp
from sympy.calculus.util import continuous_domain, function_range


def parity(num: int) -> str:
    return "Even" if num % 2 == 0 else "Odd"


def prime(num: int) -> str:
    result: int = num
    if num and True == 0:
        result = 2
    d = 3
    while d * d <= num:
        if num % d == 0:
            result = d
        d = d + 2
    
    return "Prime" if result == num else "Not Prime"


def gcd(a: int, b: int) -> int:
    r = a % b
    while r != 0:
        a, b, r = b, r, (a % b)
    return b


def prime_factors(n: int) -> str:
    factors = {}

    def loop_divisor(num, divisor):
        while num % divisor == 0:
            if divisor in factors:
                factors[divisor] += 1
            else:
                factors[divisor] = 1
            num //= divisor

        return num

    n = loop_divisor(n, 2)
    n = loop_divisor(n, 3)

    k = 5
    while n > 1 and k * k <= n:
        p1 = k
        p2 = k + 2

        n = loop_divisor(n, p1)
        n = loop_divisor(n, p2)

        k += 6

    result = ""
    if n != 1: factors[n] = 1
    for factor, count in factors.items():
        if count == 1:
            result += f"{factor} * "
        else:
            result += f"{factor}^{count} * "

    return result[:-3]


def lcm(a: int, b: int) -> int:
    return (a * b) // gcd(a, b)


def direct_proof(num: int, expr: str) -> str: # WTF?
    if expr == "If a number is even, then it is divisible by 2.":
        if parity(num) == "Even":
            result = f"Here, {num} is an even number since it can be written in the form of 2k where k={num // 2}. Hence, according to the definition, it is divisible by 2, validating the given statement."
        else:
            result = f"Here, {num} is not an even number, which violates our hypothesis as well as the conclusion"
    elif expr == "If a number is even, then it is not divisible by 2.":
        if parity(num) == "Even":
            result = f"Here, {num} is an even number since it can be written in the form of 2k where k={num // 2}. Hence, according to the definition, it is divisible by 2, which violates our hypothesis as well as the conclusion."
        else:
            result = f"Here, {num} is not an even number, which violates our hypothesis. However, it does apply to our conclusion, stating that our number {num} is not divisible by 2."

    return result


def ant_func(n: int) -> int:
    phi = int(n > 1 and n)

    for p in range(2, int(n ** 0.5) + 1):
        if not n % p:
            phi -= phi // p
            while not n % p:
                n //= p

    if n > 1:
        phi -= phi // n

    return phi


def prove_by_induction(n: int) -> None: # WTF_2?

    def sum_of_odd_numbers(n: int) -> int:
        counter = 0
        number = 1
        sum = 0

        while counter < n:
            if number % 2 == 1:
                counter += 1
                sum += number
            number += 1
        return sum

    statement = "The sum of the first n odd numbers is n^2 for all positive integers n."
    print(f"Statement: {statement}")
    print("Proof by Mathematical Induction:")

    print("\nBase Case:")
    base_result = 1
    print(f"For n = 1, the sum of the first 1 odd number is 1^2 = {base_result}")

    if base_result == 1:
        print("Base case is verified.")
    else:
        print("Base case is not verified. The statement is false.")
        return

    # Induction Step
    print(f"\nInduction Step for n = {n}:")

    # Induction Hypothesis: Assuming the statement is true for n, prove it for n + 1
    induction_hypothesis = sum_of_odd_numbers(n) == n ** 2

    if induction_hypothesis:
        print(f"The sum of the first {n} odd numbers is {n}^2 = {n ** 2}")
        print(f"Assuming the statement is true for n = {n}, the induction hypothesis is verified.")
    else:
        print(f"Assuming the statement is true for n = {n}, the induction hypothesis is not verified.")
        print("The statement is false for n =", n)
        return

    print("\nMathematical induction is successfully applied.")
    print(f"The statement is proven for all positive integers n up to {n}.")


def func_eval(f, x: sp.Symbol, x_value: int) -> str:
    domain = continuous_domain(f, x, sp.Reals)
    range_ = function_range(f, x, domain)

    result = f.subs(x, x_value)

    explanation =  f"Evaluating the function {f} at x = {x_value}:\n"
    explanation += f"f({x_value}) = {result}\n"
    explanation += f"Domain of the function: {domain}\n"
    explanation += f"Range of the function: {range_}\n"

    return explanation


if __name__ == "__main__":
    # ic(parity(25))
    #
    # ic(prime(17))
    #
    # ic(gcd(48, 18))
    #
    # ic(prime_factors(56))
    #
    # ic(lcm(15, 20))
    #
    # ic(direct_proof(10, "If a number is even, then it is not divisible by 2."))
    #
    # ic(ant_func(12))
    #
    # ic(prove_by_induction(10))
    #
    # x: sp.Symbol = sp.Symbol("x")
    # f: sp.Add = x**2 + 2*x + 1
    # ic(func_eval(f, x, 3))

    # ic(prime_factors(60))
    ic(lcm(4, 20))
    ic(prime_factors(120))
    ...

