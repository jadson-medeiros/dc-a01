import { Router } from 'express';
import pkg from 'pg';
const { Pool } = pkg;

const pool = new Pool({
  connectionString: process.env.POSTGRES
});

const itemsRouter = Router();

itemsRouter.get('/', async (req, res) => {
  try {
    sql = 'SELECT * FROM tb01 ORDER BY col_dt DESC limit 10'
    const response = await pool.query(sql);
    res.status(200).send(response.rows);
  } catch (err) {
    return res.status(500).send(err);
  }
});

export default itemsRouter;