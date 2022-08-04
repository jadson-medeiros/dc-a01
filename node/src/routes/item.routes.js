import express from 'express';
const router = express.Router({ mergeParams: true });

import pkg from 'pg';
const { Pool } = pkg;

const pool = new Pool({
  connectionString: "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
});

router.get('/', async (req, res) => {
  try {
      const { rows } = await pool.query('SELECT * FROM TB01 ORDER BY col_dt DESC LIMIT 10')
      return res.status(200).send(rows)
  } catch (error) {
      return res.status(400).send(error)
  }
})

export default router;