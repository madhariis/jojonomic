import express from 'express'
import cors from 'cors'
import mongoose from 'mongoose'
import router from'./router'
const dotenv = require('dotenv')

dotenv.config()

const app = express()
const port = process.env.PORT || 3000

mongoose.connect('mongodb://localhost:27017/document-service', {useNewUrlParser: true, useUnifiedTopology: true})

app.use(cors())
app.use(express.urlencoded({extended: true}))
app.use(express.json())
app.use(router)


app.listen(port, () => {
    // console.log(process.env)
    console.log(`listen to port: ${port}`)
})