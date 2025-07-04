import {BrowserRouter, Routes, Route} from 'react-router-dom';
import { BeforeShortened } from './BeforeShortened';
import { AfterShortened } from './AfterShortened';
export default function PageIndex(){
    return(
        <div>
            <BrowserRouter>
                <Routes>
                    <Route path='/' element={<BeforeShortened/>} />
                    <Route path='/shortener' element={<AfterShortened/>} />
                </Routes>
            </BrowserRouter>
        </div>
    )
}



