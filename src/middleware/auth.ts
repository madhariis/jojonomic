import verifyToken from '../helper/jwt'

function authorization(req, res, next){
    let { token } = req.headers
    try {
        let decoded = verifyToken(token)
        if(decoded){
            let { user_id, company_id } = decoded
            req.user_id = user_id
            req.company_id = company_id
            next()
        } else {
            res.status(401).json({
                message: 'not authorized'
            })
        }
    } catch (error) {
        console.log(error)
        res.status(500).json({
            message: error.message
        })
    }
}

export default authorization