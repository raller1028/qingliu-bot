from flask import Flask
import requests
import base64
import urllib.parse
from bs4 import BeautifulSoup
from Cryptodome.Util.Padding import pad
from Cryptodome.Cipher import AES

app = Flask(__name__)


# encrypts the password using the public key

def encrypt(key, password):
    aes = AES.new(key.encode('utf-8'), AES.MODE_ECB)
    pad_pkcs7 = pad(password.encode('utf-8'), AES.block_size, style='pkcs7')
    encrypt_aes = aes.encrypt(pad_pkcs7)
    encrypted_text = str(base64.encodebytes(encrypt_aes), encoding='utf-8')  # 解码
    encrypted_text_str = encrypted_text.replace("\n", "")
    return encrypted_text_str


@app.route('/')
def hello_world():
    session = requests.Session()

    # 1 get key,execution,sesssion,encrypted password from login page
    loginPageRes = requests.get('https://sso.4px.com/login?service=https://b.4px.com/cas&amp;from=bcp')
    soup = BeautifulSoup(loginPageRes.text, 'lxml')
    execution = urllib.parse.quote(soup.find('input', {'name': 'execution'})['value'])
    key = soup.find('input', {'name': 'key'})['value']
    password = urllib.parse.quote(encrypt(key, 'Icewhale2020'))
    loginPageCookies = loginPageRes.cookies.get_dict()['SESSION']

    # print(execution)
    # print(key)
    # print(loginPageCookies)
    # print(password)

    # 2 sso login
    url = "https://sso.4px.com/login?service=http://b.4px.com/cas&from=bcp"
    payload = 'username=13818662085&imgverifycode=&lt=&execution=' + execution + '&_eventId=submit&key=' + key + '&password=' + password
    headers = {
        'authority': 'sso.4px.com',
        'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9',
        'accept-language': 'zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7',
        'cache-control': 'max-age=0',
        'content-type': 'application/x-www-form-urlencoded',
        'cookie': 'SESSION=' + loginPageCookies,
        'origin': 'https://sso.4px.com',
        'referer': 'https://sso.4px.com/login?service=https://b.4px.com/cas&amp;from=bcp',
        'sec-ch-ua': '"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
        'sec-fetch-dest': 'document',
        'sec-fetch-mode': 'navigate',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-user': '?1',
        'upgrade-insecure-requests': '1',
        'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36'
    }

    response = session.request("POST", url, headers=headers, data=payload)

    # print(response.text)

    # 3 get ticket
    payload1 = "service=https%3A%2F%2Forder.4px.com%2Fcas"

    headers1 = {
        'content-type': 'application/x-www-form-urlencoded;charset=UTF-8',
    }
    response1 = session.request("POST", "https://sso.4px.com/v1/tickets/" + session.cookies.get_dict()['tgt_cn'],
                                headers=headers1, data=payload1)

    ticket = response1.text

    # 4 login to order.4px.com
    url2 = "https://order.4px.com/cas?ticket=" + ticket

    headers2 = {
        'x-requested-with': 'XMLHttpRequest',
    }

    response3 = session.request("GET", url2, headers=headers2)

    print(response3.text)
    Session = response3.cookies.get_dict()["SESSION"]

    return Session


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
