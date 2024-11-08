import { useState } from 'react';
import './Form.css';

export default function URLForm() {
    const [longUrl, setLongUrl] = useState('');
    const [alias, setAlias] = useState('');

    const handleSubmit = (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        // Perform your submit logic here, e.g., API request
        console.log('Long URL:', longUrl);
        console.log('Alias:', alias);
        // Reset form fields after submission
        setLongUrl('');
        setAlias('');
    };
    return (
        <>
            <div className='virtual-body'>
                <div className="url-form">
                    <form onSubmit={handleSubmit}>
                        <div className='layer'>
                            <label htmlFor="long-url">Shorten a long URL:</label>
                            <input
                                type="text"
                                id="long-url"
                                placeholder="Enter a long URL"
                                value={longUrl}
                                onChange={(event) => setLongUrl(event.target.value)}
                                required />
                        </div>
                        <div className='layer'>
                            <label htmlFor="alias">Customize your link:</label>
                            <input type="text"
                                id="alias"
                                placeholder="Enter Alias (optional)"
                                value={alias}
                                onChange={(event) => setAlias(event.target.value)} />
                        </div>
                        <button type="submit" id='submit'>Shorten URL</button>
                    </form>
                </div>
            </div>
        </>
    )
}