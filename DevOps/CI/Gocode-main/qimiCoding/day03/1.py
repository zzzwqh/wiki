import pandas as pd
# 读取Excel文件
df = pd.read_excel(r'D:\Users\ethan\GolandProjects\GoCode\qimi\day03\网段.xlsx')
print(df)
 # 分组规则
subnet_ranges = {
    'UAT网段': '192.168.1.',
    'DR环境': '10.120.0.',
    'DEV': '10.121.1.',
    'PROD环境': '172.161.1.',
    'Azure测试环境':'10.0.0.',
    'Azure生产环境':'10.130.0.',
    'AzureDR环境':'10.122.2.',
    'AzureINFRA环境':'172.162.1.'
}
#print (subnet_ranges)
# 循环遍历查找IP数据
for subnet_name, subnet_range in subnet_ranges.items():
# 构建筛选条件
    conditions = []
    for subnet in subnet_range:
        condition = df['IP'].str.startswith('subnet')
        conditions.append(conditions)
        # print(conditions)
        str = pd.concat(conditions, axis=0).any()
        # 筛选出特定网段的数据
        subnet_df = df[str]
        # # 将数据写入Excel文件
