import React, { useEffect,useState } from 'react'
import { useSnackbar  } from 'notistack';
import SnackbarUtils from './SnackbarUtils';
import { useHistory } from "react-router-dom";


const Snackbar = () => {
    const { enqueueSnackbar,closeSnackbar } = useSnackbar();
    const [urls,setUrls] = useState([])


    const handleClick = (text) => {
        SnackbarUtils.info(text);
      };

    useEffect(() => {
       setUrls()

        SnackbarUtils.setSnackBar(enqueueSnackbar,closeSnackbar)

    }, [])

    return (
        <div>
            
        </div>
    )
}

export default Snackbar