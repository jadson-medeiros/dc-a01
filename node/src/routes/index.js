import { Router } from 'express';
import itemsRouter from './item.routes.js';

const routes = Router();

routes.use('/tb01', itemsRouter);

export default routes;