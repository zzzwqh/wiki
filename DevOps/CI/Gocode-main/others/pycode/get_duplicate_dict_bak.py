input_str = input()
position = {}


# 输入一段字符串，第一个重复的字符打印出来
def get_first_duplicate(arg):
    index = 0
    for i in arg:
        if i not in position.keys():
            position[i] = arg.index(i)
        else:
            arg
            char_dict[i] = index
            index = index + 1


print(get_first_duplicate(input_str))
print(char_dict)
