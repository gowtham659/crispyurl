import {useFormik} from 'formik';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useState } from 'react';
import './styles/beforestyles.css'


export function BeforeShortened(){
    const navigate = useNavigate()
    const [reject,setReject] =  useState("")

    //const currentUrl = "http://localhost:8989/";
    const currentUrl = "https://react-go-oracle-app.onrender.com/";
    const formik = useFormik({
        initialValues:{Url:""},
        onSubmit: (values)=>{
            axios.post(currentUrl+"shorten",{
                "OrgUrl": values.Url,
                "BaseUrl": currentUrl
            })
            .then((values)=>{
                let responseData = values.data;
                console.log(responseData)
                if (responseData.message ==="success"){
                    console.log("navigating to /shortener ")
                    navigate("/shortener",{state: {"shorturl":responseData["shorten_url"]}})
                }
                else{
                    setReject("Unable to generate Short URL!!!")
                }
            })
            .catch(error=>{
                console.log(error.message, error.stack)
            })
            
        },
        
    })
    return(
        <div>
            <div className="container">
                <div className="title">
                    Short URL
                </div>
                <div className="row">
                    <div className="row-title">
                        Paste the URL to be shortened
                    </div>
                    <div className="row-input-button">
                        <form onSubmit={formik.handleSubmit}>
                            <span className="row-input">
                                <input type="text" name="Url" placeholder="Enter the link here" onChange={formik.handleChange} required/>
                            </span>
                            <span className="row-button"><button type='submit' >Shorten URL</button></span>
                        </form>
                        <div className='row-reject'>
                            <p>{reject}</p>
                        </div>
                    </div>
                    <div className="row-description">
                        <div className='row-description-1'>ShortURL is a free tool to shorten URLs and generate short links</div>
                        <div className='row-description-2'>URL shortener allows to create a shortened link making it easy to share</div>
                    </div>
                </div>
            </div>
        </div>
    )
}
