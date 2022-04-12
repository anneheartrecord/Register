'''
我们这里用 1 2 3 模拟操作 1 代表注册 2 代表登录 3 代表退出系统
很容易想到的是 账号密码使用k-v键值对存储
'''
dictionary={}
def check(username,password):
    if username in sorted(dictionary.keys()):
        print("用户名已被注册")
        return False
    if len(password)<=6:
        print("密码不能小于6位")
        return False
    return True

def Register():
    print("请输入你的账号")
    username=input()
    print("请输入你的密码")
    password=input()
    if check(username,password):
        dictionary[username]=password
        print("注册成功")
def Login():
    print("请输入你的账号")
    username=input()
    print("请输入你的密码")
    password=input()
    if password==dictionary[username]:
        print("登陆成功")
    else:
        print("密码错误")



print("1.用户注册 2.用户登录 3.退出系统")
print("请输入你的操作")
opt=input()
while True:
    if opt=='1':
        Register()
    if opt=='2':
        Login()
    if opt=='3':
        break
    opt=input("请输入你的操作\n")
