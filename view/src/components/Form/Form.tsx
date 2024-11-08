import './Form.css';

export default function URLForm() {
    return (
        <>
            <div className='virtual-body'>
                <div className="url-form">
                    <form>
                        <div className='layer'>
                            <label htmlFor="long-url">Shorten a long URL:</label>
                            <input type="text" id="long-url" placeholder="Enter a long URL" />
                        </div>
                        <div className='layer'>
                            <label htmlFor="alias">Customize your link:</label>
                            <input type="text" id="alias" placeholder="Enter Alias" />
                        </div>
                        <input type="submit" id="submit" value="Shorten URL" />
                    </form>
                </div>
            </div>
        </>
    )
}