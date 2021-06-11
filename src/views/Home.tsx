const Home = () => {
    const gifs = [
        "https://media1.tenor.com/images/47fcadc56b78a010732a75e8185c3352/tenor.gif?itemid=9950777",
        "https://media1.tenor.com/images/382f7096ff1e0c01a8124824f3d6d916/tenor.gif?itemid=6318866",
        "https://media1.tenor.com/images/7853f41dac5f6f50c0c98488013838d4/tenor.gif?itemid=13755444",
        "https://media1.tenor.com/images/dcc032817ab75342c947365051513701/tenor.gif?itemid=17047532",
        "https://media1.tenor.com/images/a3499cebad66f3cf8194b646a95b38ba/tenor.gif?itemid=16284566",
        "https://media1.tenor.com/images/399c44e34035c563aa09488724d5df28/tenor.gif?itemid=16865940",
        "https://i.giphy.com/media/ChzovjKPuEiYe8ePih/giphy.webp",
        "https://media.giphy.com/media/3o85xu3OLHOVvzZNra/giphy.gif",
        "https://media.giphy.com/media/l1Et6k00qp9fMTP8s/giphy.gif",
        "https://media.giphy.com/media/26grMgCg1xZh28AF2/giphy.gif"

    ]
    const randGif = Math.floor(Math.random() * gifs.length)

    return (
        <div className="home">
            <h1>Welcome!</h1>
            <img alt="" src={gifs[randGif]} />
        </div>
    )
}

export default Home