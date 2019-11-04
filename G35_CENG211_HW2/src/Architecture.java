
public class Architecture extends Artworks {
	private int dimension1;
	private int dimension2;
	private int dimension3;
	private Architect architect;

	public Architecture(String name,String style,int dimension1,int dimension2,int dimension3,Architect architect) {
		super(name,style);
		this.dimension1 = dimension1;
		this.dimension2 = dimension2;
		this.dimension3 = dimension3;
		this.architect = architect;
	}
}
