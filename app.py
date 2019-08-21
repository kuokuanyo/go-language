#LINE聊天機器人
#flask架設伺服器或網站
#flask架設規模較小，django規模較大
#github做版本控制
#heroku執行程式碼
from flask import Flask, request, abort

from linebot import (
    LineBotApi, WebhookHandler
)
from linebot.exceptions import (
    InvalidSignatureError
)
from linebot.models import (
    MessageEvent, TextMessage, TextSendMessage, StickerSendMessage
)

app = Flask(__name__)

line_bot_api = LineBotApi('FNsM9L8athPD76AwtTprDyKAbyxA8jL2FbZyIzdjIkNPluoZKXoYq18few+jj0ylVDEMX9tYNhVuPYoQd1bWcq60oXk5uH7BjQU/rcGMBS53p5vNvRGYPctwWQDJ0vtVhPsvK2Wa5XS2Y9bOuWi0EQdB04t89/1O/w1cDnyilFU=')
handler = WebhookHandler('e1b38a26e2d3a9e038b0d8c87cde296c')


@app.route("/callback", methods=['POST'])
def callback():
    # get X-Line-Signature header value
    signature = request.headers['X-Line-Signature']

    # get request body as text
    body = request.get_data(as_text=True)
    app.logger.info("Request body: " + body)

    # handle webhook body
    try:
        handler.handle(body, signature)
    except InvalidSignatureError:
        print("Invalid signature. Please check your channel access token/channel secret.")
        abort(400)

    return 'OK'


@handler.add(MessageEvent, message=TextMessage)
def handle_message(event):
    msg = event.message.text
    r = '我看不懂你說什麼'

    if '貼圖' in msg:
        sticker_message = StickerSendMessage(
            package_id='1',
            sticker_id='1'
        )
        
        line_bot_api.reply_message(
            event.reply_token,
            sticker_message)
        return

    if msg in ['hi','Hi']:
        r = '哈囉'
    elif msg == '你吃飯了嗎':
        r = '還沒'
    elif msg == '你是誰':
        r = '我是機器人'
    elif '訂位' in msg:
        r = '您想訂位，是嘛?'
    line_bot_api.reply_message(
        event.reply_token,
        TextSendMessage(text=r))


if __name__ == "__main__":
    app.run()