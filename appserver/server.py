from fastapi import FastAPI,logger

from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()
app.add_middleware(CORSMiddleware,allow_origins=["*"],allow_methods=["*"],allow_headers=["*"])


@app.get('/enter')
def firstfxn():
    return {"here":"akwaaba"}
