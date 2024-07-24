import sys
import requests

# 检查命令行参数
if len(sys.argv) < 4:
    print("Usage: python script.py <host:port> <file_path> <access_password>")
    sys.exit(1)

# 获取主机、文件路径和访问密码
host_port = sys.argv[1]
file_path = sys.argv[2]
access_password = sys.argv[3]

# 构建URL
url = f'http://{host_port}/'

# 请求头中的Cookie
headers = {
    'Cookie': f'access_password={access_password}'
}

# 使用requests的post方法上传文件
try:
    with open(file_path, 'rb') as file:
        files = {'file': (file_path, file)}
        response = requests.post(url, headers=headers, files=files)

    # 打印响应内容
    print(response.text)
    print(f"-----------------------")
    print(f"Url:http://" + host_port + "/" + file_path)
except FileNotFoundError:
    print(f"Error: The file '{file_path}' does not exist.")
except Exception as e:
    print(f"An error occurred: {e}")
