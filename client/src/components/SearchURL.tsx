import React, { useState } from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import GetAppIcon from '@material-ui/icons/GetApp';
import CircularProgress from '@material-ui/core/CircularProgress';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

const BACKEND_BASE_URL = "http://localhost:8080"

const useStyles = makeStyles((theme: Theme) => 
  createStyles({  
    root: {
      '& > *': {
        margin: 'auto',
        width: '70ch',
      },
    },
    button: {
      width: '10rem',
      height: '2rem',
      boxShadow: 'none',
      marginTop: '20px',
      padding: '7px 14px',
      backgroundColor: '#242526',
      color: '#ededed',
      borderRadius: '5px',
      textTransform: 'none'
    },
    iconButton: {
      marginTop: '20px',
      color: '#ededed',
      width: '50px',
      borderRadius: '5px'
    }
  }),
)

const SearchURL = () => {
  const [sourceUri, setSourceUri] = useState('');
  const [inputUri, setInputUri] = useState('https://karngyan.com');
  const [helperText, setHelperText] = useState('Please enter the URL')
  const [downloadDisabled, setDownloadDisabled] = useState(true);
  const [loading, setLoading] = useState(false);

  const classes = useStyles();

  async function requestPageSource() {
    setLoading(true);
    const jsonHeaders = new Headers();
    jsonHeaders.append('Content-Type', 'application/json');

    const data = {
      uri: inputUri,
      retryLimit: 4
    };

    const requestOptions = {
      method: 'POST',
      headers: jsonHeaders,
      body: JSON.stringify(data),
    };

    const endpoint = BACKEND_BASE_URL + '/pagesource';

    const response = await fetch(endpoint, requestOptions);
    const responseJson = await response.json();
    const newSourceUri = BACKEND_BASE_URL + responseJson.sourceUri;
    setSourceUri(newSourceUri);
    setTimeout(() => {
      setDownloadDisabled(false);
      setLoading(false);
      setHelperText("Expected Path URI: " + newSourceUri)
    }, 1500);
  }

  const renderDownloadButton = () => {
    return (
      <IconButton 
        className={classes.iconButton} 
        color="primary" 
        aria-label="download html file" 
        disabled={downloadDisabled} 
        href={sourceUri}
        target="_blank">
        <GetAppIcon />
      </IconButton>
    )
  }

  return <div className={classes.root}>
    <form autoComplete="off" onSubmit={(e) => {
      e.preventDefault();
      requestPageSource();
    }}>
      <TextField
      value={inputUri}
      helperText={helperText}
      onChange={(e) => setInputUri(e.target.value)}
      id="outlined-basic" 
      label="URL" 
      variant="outlined"
      required
      fullWidth
      ></TextField>
      <Button className={classes.button} size="small" type="submit" disableRipple>
        {loading ? <CircularProgress color="primary" size="1.5rem"/> : "Get Page Source"}
      </Button>
    </form>
    {renderDownloadButton()}
  </div>
}

export default SearchURL;