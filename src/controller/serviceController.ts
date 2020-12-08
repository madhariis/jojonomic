import Document from '../model/documentModel'
import Folder from '../model/folderModel'
import redis from '../redis'

class DocumentController{
    static async getAll(req, res){ 
        try {
            const cache = await redis.get('services')
            if(cache){
                res.status(200).json({
                    error: false,
                    data: JSON.parse(cache)}) 
            } else {
                let data : object[] = []
                let documents = await Document.find()
                let folders = await Folder.find()
                data = documents.concat(folders)
                await redis.set('services', JSON.stringify(data))
                res.status(200).json({
                    error: false,
                    data
                })
            }
        } catch (error) {
            res.status(500).json({error})
        }
    }

    static async getListFile(req, res){
        let { folder_id } = req.params
        try {
            const cache = await redis.get('folder')
            if(cache){
                let data = JSON.parse(cache)
                let isSame = (value) => value == folder_id
                let datas = data.map(elm => {
                               return elm.folder_id
                }) 
                if(datas.every(isSame)){
                    res.status(200).json({
                        error: false,
                        data
                    })
                } else {
                    data = await Document.find({folder_id})
                    await redis.set('folder', JSON.stringify(data))
                    res.status(200).json({
                    error: false,
                    data
                })    
                }
            } else {
                const data = await Document.find({folder_id})
                await redis.set('folder', JSON.stringify(data))
                res.status(200).json({
                    error: false,
                    data
                })
            }
        } catch (error) {
            res.status(500).json({
                error
            })
        }
    }

    static async getDetailDocument(req, res){
        let { document_id } = req.params
        try {
            const cache = await redis.get('document')
            if(cache){
                let data = JSON.parse(cache)
                if(data.id == document_id){
                    res.status(200).json({
                        error: false,
                        data
                    })    
                } else {
                    data = await Document.findOne({id: document_id})
                    await redis.set('document', JSON.stringify(data))
                    res.status(200).json({
                        error: false,
                        data
                    })        
                }
            } else {
                const data = await Document.findOne({id: document_id})
                await redis.set('document', JSON.stringify(data))
                res.status(200).json({
                    error: false,
                    data
                })
            }
        } catch (error) {
            res.status(500).json({
                error
            })
        }
    }

    static async setFolder(req, res){
        let {id, name, timestamp} = req.body
        let user_id = req.user_id
        let company_id = req.company_id
        try {
            const checker = await Folder.findOne({id})
            if(checker){
                let folder = ({
                    id,
                    name,
                    timestamp
                })
                let data = await Folder.replaceOne(folder)
                res.status(201).json({
                    error: false,
                    message: 'folder updated',
                    data
                })
            } else {
                let folder = new Folder({
                    id,
                    name,
                    timestamp,
                    owner_id: user_id,
                    company_id
                })
                let cache = await redis.get('services')
                let cacheData = JSON.parse(cache)
                let data = await folder.save()
                cacheData.push(data)
                await redis.set('services', JSON.stringify(cacheData))
                res.status(201).json({
                    error: false,
                    message: 'folder created',
                    data
                })
            }
        } catch (error) {
            res.status(500).json({
                error
            })
        }
    }

    static async setDocument(req, res){
        let { id, name, type, folder_id, content, timestamp, share, company_id } = req.body
        try {
            const newDocument = new Document({
                id,
                name,
                type,
                folder_id,
                content,
                timestamp,
                share,
                company_id
            })
            const cache = await redis.get('services')
            const cacheData = JSON.parse(cache)
            const document = await newDocument.save()
            cacheData.push(document)
            await redis.set('services', JSON.stringify(cacheData))
            res.status(201).json({
                error: false,
                message: 'success set document',
                data: document
            })
        } catch (error) {
            res.status(500).json({
                error
            })
        }
    }

    static async deleteFolder(req, res){
        let { id } = req.body
        try {
            let cache = await redis.get('services')
            let cacheData = JSON.parse(cache)
            for (let i = 0; i < cacheData.length; i++) {
                if(cacheData[i].id == id){
                    cacheData.splice(i, 1)
                }
            }
            await redis.set('services', JSON.stringify(cacheData))
            let data = await Folder.findOneAndDelete({id})
            res.status(200).json({
                error: false,
                message: 'Success delete folder'
            })
        } catch (error) {
            res.status(500).json({
                error: 'internal server error'
            })
        }
    }

    static async deleteDocument(req, res){
        let { id } = req.body
        try {
            let cache = await redis.get('services')
            let cacheData = JSON.parse(cache)
            for (let i = 0; i < cacheData.length; i++) {
                if(cacheData[i].id == id){
                    cacheData.splice(i, 1)
                }
            }
            await redis.set('services', JSON.stringify(cacheData))
            let data = await Document.findOneAndDelete({id})
            res.status(200).json({
                error: false,
                message: 'Success delete document'
            })
        } catch (error) {
            res.status(500).json({
                error: 'internal server error'
            })
        }
    }
}

export default DocumentController