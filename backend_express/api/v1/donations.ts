import { Router, Request, Response } from 'express';
import { getDonations } from './repositories/donations';

const router = Router();

router.get('/', (req: Request, res: Response) => {
    const donations = getDonations();
    res.json(donations);
});

export default router;
