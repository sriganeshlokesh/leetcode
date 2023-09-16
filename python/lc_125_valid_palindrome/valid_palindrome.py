def valid_palindrome(s: str) -> bool:
    i, j = 0, len(s) - 1

    while i < j:
        while i < j and s[i].isalnum() == False:
            i += 1
        while j > i and s[j].isalnum() == False:
            j -= 1
        if i > j or s[i].lower() != s[j].lower():
            return False
        else:
            i += 1
            j -= 1

    return True

if __name__ == "__main__":
    print(valid_palindrome("A man, a plan, a canal: Panama"))