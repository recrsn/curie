import React from 'react';
import './URIMaker.scss';

export default class URIMaker extends React.Component {
    render() {
        return (
            <div className="URIMaker">
                <h1>
                    Simplify your links
                </h1>
                <form className="pure-form">
                    <div className="row">
                        <input className="url-field pure-input-1" type="url" placeholder="Your original URI here" />
                        <button className="btn-create pure-button" type="submit">Shorten URL</button>
                    </div>
                    <div>
                        {/* <label>
                            <input type="checkbox" defaultChecked={true} />
                            <span>Auto generate short link</span>
                        </label> */}
                    </div>
                </form>
            </div>
        );
    }
}
