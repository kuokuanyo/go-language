import React, { Component } from 'react';
import { books } from '../data';

class Books extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div>
                <table className="table table-striped">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Title</th>
                            <th>Author</th>
                            <th>Year</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            books.map(book => {
                                return (
                                    <tr key={book.id}>
                                        <td>{book.id}</td>
                                        <td>{book.title}</td>
                                        <td>{book.author}</td>
                                        <td>{book.year}</td>
                                        <td>
                                            <button type="button" className="btn btn-danger">
                                                Delete
                                            </button>
                                            <button type="button" className="btn btn-danger">
                                                Edit
                                            </button>                                           
                                        </td>
                                    </tr>
                                )
                            })
                        }
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Books;
