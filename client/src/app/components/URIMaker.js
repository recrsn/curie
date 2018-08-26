import React from 'react';
import './URIMaker.css';

export default class URIMaker extends React.Component {
    render() {
        return (
            <iv className="URIMaker">
                <h1>
                    Simplify your links
                </h1>
                <form className="pure-form">
                    <div className="row">
                        <input className="url-field" type="url" placeholder="Your original URI here" />
                        <button className="pure-button pure-button-primary btn-create" type="submit">Shorten URL</button>
                    </div>
                    <div>
                        {/* <label>
                            <input type="checkbox" defaultChecked={true} />
                            <span>Auto generate short link</span>
                        </label> */}
                    </div>
                </form>
            </iv>
        );
    }
}
