import { useLocation } from "react-router-dom"
import './styles/afterstyles.css'
import { useState } from "react"



export function AfterShortened(){
    const location = useLocation()
    const url = location.state
    const [copy, setCopy] = useState("Copy")
    console.log(url)
    return(
        <div>
            <div className='container-2'>
                <div className='title'>
                    <div className='title-1'>Your shortened URL</div>
                    <div className='title-2'>Copy the short link and share it in messages, texts, posts, websites and other locations.</div>
                </div>
                <div className='row' >
                    <div className="row-input-button">
                        {/* <input type="text" name="Url" placeholder="Enter the link here" value={url.shorturl} /> */}
                        <div>
                            <input
                            type="text"
                            name="Url"
                            value={url.shortUrl}
                            readOnly
                            onClick={(e) => e.target.select()}
                        />
                        <button onClick={() => {
                            navigator.clipboard.writeText(url.shortUrl)
                            setCopy("Copied")
                        }}>{copy}</button>
                        </div>
                        <div className="other-btn">
                            <a href={url.shortUrl} target="_blank" rel="noreferrer"><button  id="btn1"><li class="fa fa-external-link"></li></button></a>
                            <button id="btn2"><a href={url.appUrl}  >shorten another URL</a></button>
                        </div>
                        
                    </div>
                </div>
            </div>
        </div>
    )
}