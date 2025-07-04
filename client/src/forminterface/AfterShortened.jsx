import { useLocation } from "react-router-dom"
import './styles/afterstyles.css'
export function AfterShortened(){
    const location = useLocation()
    const url = location.state
    console.log(url)
    return(
        <div>
            <div className='container-2'>
                <div className='title'>
                    <div className='title-1'>Your shortened URL</div>
                    <div className='title-2'>Copy the short link and share it in messages, texts, posts, websites and other locations.</div>
                </div>
                <div className='row'>
                    <div className="row-input-button">
                        {/* <input type="text" name="Url" placeholder="Enter the link here" value={url.shorturl} /> */}
                        <input
                            type="text"
                            name="Url"
                            value={url.shorturl}
                            readOnly
                            onClick={(e) => e.target.select()}
                        />
                        <button onClick={() => navigator.clipboard.writeText(url.shorturl)}>Copy</button>
                    </div>
                </div>
            </div>
        </div>
    )
}