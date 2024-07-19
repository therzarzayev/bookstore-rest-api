import os

user_list = ['alpha', 'beta', 'gamma']
group = "devs"

gexitcode = os.system("grep {} /etc/group >/dev/null".format(group))
if gexitcode != 0:
    os.system("sudo groupadd {}".format(group))
else:
    print("Group already exist!")


for user in user_list:
    uexitcode = os.system("grep {} /etc/passwd >/dev/null".format(user))
    if uexitcode != 0 :
        os.system("sudo useradd {}".format(user))
    else:
        print("'{}' user exists on system".format(user))
    os.system("sudo usermod -aG {} {} > /dev/null".format(group, user))
    print("User added to group")

