import express from 'express';

const app = express();

const tb01 = [
  {id: 1, "col_texto":"test", "col_dt": Date.now()},
]

app.get('/tb01', (req, res) => {
  res.status(200).json(tb01);
})

export default app;