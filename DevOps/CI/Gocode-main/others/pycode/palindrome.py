input_int = int(input())
# 3 3210123
# 2 21012
def palindrome(num):
    if num == 0:
        return 0
    else:
        left = num*10**(2*num)
        for i in range(1,num):
            left += (num-i)*10**(2*num-i)
        mid = palindrome(0)
        right = num
        for i in range(num,1,-1):
            right += (i-1)*10**(num-i+1)
        return left + mid + right


print(palindrome(input_int))