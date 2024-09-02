input_str = input()
char_set = set()


# 输入一段字符串，第一个重复的字符打印出来
def get_first_duplicate(arg):
    for i in arg:
        if i in char_set:
            print(arg.index(i))
            return i
        else:
            char_set.add(i)


print(get_first_duplicate(input_str))
