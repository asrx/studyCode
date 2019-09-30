import React from 'react'

class HelloWorld extends React.Component {
    render() {
        return (
            <div className="hello">
                <h1>{this.props.msg}</h1>
            </div>
        )
    }

}

export default HelloWorld
