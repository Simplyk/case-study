import { Router, Request, Response } from 'express';
import * as donations from "./100-last-donations.json";

const router = Router();

router.get('/', (req: Request, res: Response) => {
    res.json(donations);
});

export default router;
