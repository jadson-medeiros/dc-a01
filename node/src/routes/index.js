import express from 'express';
import itemsRouter from './item.routes.js';

const routes = express.Router();

routes.use('/tb01', itemsRouter);

export default routes;