FROM python:3.8-alpine

RUN python3 -m pip install --upgrade pip

WORKDIR /app
COPY requirements.txt .
RUN pip3 install -r requirements.txt
COPY main.py .
CMD ["python3", "-u", "main.py"]