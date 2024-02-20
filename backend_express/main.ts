import * as express from "express";
import donationsApi from "./api/v1/donations"
import * as cors from "cors";

const app = express();
const port = 1323;

app.use(express.json());
app.use(cors());

app.use("/api/v1/donations", donationsApi);

app.get('/', (req, res) => {
  res.send('Hello World!')
});

app.listen(port, () => {
  console.log(`App listening on port ${port}`)
});
