import "./Navbar.css" 


const Navbar = () => {
    return (
        <header>
            <section>
                <ul>
                        <li>
                    <a href="#Posts">
                            Posts
                    </a>
                        </li>
                    <a href="#Contratantes">
                        <li>
                            Contratantes
                        </li>
                    </a>
                    <a href="#Freelancers">
                        <li>
                            Freelancers
                        </li>
                    </a>
                </ul>
            <img src="Gojo.png" alt="Perfil" srcset="" width={70} height={70}/>
            </section>

        </header>
    )
}

export default Navbar