import { Card, Grid, Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import React from 'react';
import './App.css';
import { Transactions } from './components';

const useStyles = makeStyles((theme) => ({
  root: {
    padding: theme.spacing(4)
  },
  titleText: {
    paddingBottom: theme.spacing(2),
    paddingTop: theme.spacing(2),
  },
}));

function App() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid
        container
        spacing={4}
      >
        <Grid
          item
          xs={12}
        >
          <Card
            className={classes.root}
          >
            <div>
              <Typography variant="h2">
                Transactions
              </Typography>
            </div>
          </Card>
        </Grid>
        <Grid
          item
          xs={12}
        >
          <Transactions />
        </Grid>
      </Grid>

    </div >
  );
}

export default App;
