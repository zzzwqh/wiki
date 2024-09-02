input_str = input()
char_dict = {}


# 输入一段字符串，第一个重复的字符打印出来
def get_first_duplicate(arg):
    index = 0
    for i in arg:
        if i in char_dict.keys():
            print(char_dict.keys())
            print("重复字符第一次出现的 index :",arg.index(i))
            print("重复字符第二次出现的 index :",index)
            return i
        else:
            char_dict[i] = index
            index = index + 1


print(get_first_duplicate(input_str))
print(char_dict)
