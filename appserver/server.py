from fastapi import FastAPI,Request,logger
import uvicorn

from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()
app.add_middleware(CORSMiddleware,allow_origins=["*"],allow_methods=["*"],allow_headers=["*"])


@app.get('/enter')
def firstfxn():
    return {"here":"akwaaba"}



if __name__ =="__main__":
    uvicorn.run(app,reload = True, port= 8000, host="0.0.0.0")    
