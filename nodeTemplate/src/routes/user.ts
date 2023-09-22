import { Router, Response } from "express";

const router: Router = Router();

router.get('/', async (req: any, res: Response) => {
    try {
    return res.send('This is the user api').status(200);
    } catch (e) {
        console.log(e);
    }
})

export default router;