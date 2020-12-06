import { Router } from 'express'
import serviceController from '../controller/serviceController'
import authorization from '../middleware/auth'

const route = Router()
route.use(authorization)
route.get('/document-service', serviceController.getAll)
route.get('/document-service/folder/:folder_id', serviceController.getListFile)
route.get('/document-service/document/:document_id', serviceController.getDetailDocument)
route.post('/document-service/folder', serviceController.setFolder)
route.post('/document-service/document', serviceController.setDocument)
route.delete('/document-service/folder', serviceController.deleteFolder)
route.delete('/document-service/document', serviceController.deleteDocument)

export default route