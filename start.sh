#!/bin/bash
python /app/qingliuSession/getSession.py &
/go/bin/lark-bot &
wait
