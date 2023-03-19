alphabet_lower = "abcdefghijklmnopqrstuvwxyz"
ALPHABET_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

def encrypt(n, plaintext):
    #Encrypt string and return ciphertext
    result = ''
    for letter in plaintext:
        try:
            if letter.isupper():
                index = ALPHABET_upper.index(letter)
                i = (index + n) % 26
                result += ALPHABET_upper[i]
            else:
                index = alphabet_lower.index(letter)
                i = (index + n) % 26
                result += alphabet_lower[i]
        except ValueError:
            result += letter
    return result

def decrypt(n, ciphertext):
    #Decrypt string and return plaintext
    result = ''
    for letter in ciphertext:
        try:
            if letter.isupper():
                index = ALPHABET_upper.index(letter)
                i = (index - n) % 26
                result += ALPHABET_upper[i]
            else:
                index = alphabet_lower.index(letter)
                i = (index - n) % 26
                result += alphabet_lower[i]
        except ValueError:
            result += letter
    return result

shift = 3
message = str(input("Enter a message: " ))
enc = encrypt(shift, message)
print("%d . %s" % (shift, enc))

dec = decrypt(shift, enc)
print("%d . %s" % (shift, dec))


